package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/application"
	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/domain/entities"
	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type Resp struct {
	Courses []entities.Course `json:"Courses"`
	Persons []entities.Person `json:"Persons"`
}

func Parse(w http.ResponseWriter, r *http.Request) {
	log := logger.GetFromContext(r.Context())
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Error(r.Context(), "Unable to extract file",
			zap.Error(err))
		http.Error(w, "Unable to extract file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	f, err := excelize.OpenReader(file)
	if err != nil {
		log.Error(r.Context(), "Unable to open excel file",
			zap.Error(err))
		http.Error(w, "Unable to extract file", http.StatusInternalServerError)
		return
	}

	sheets := f.GetSheetList()
	rowsElectives, err := f.GetRows(sheets[0])
	if err != nil {
		http.Error(w, "Failed to read first sheet", http.StatusInternalServerError)
		return
	}
	rowsCourses, err := f.GetRows(sheets[1])
	if err != nil {
		http.Error(w, "Failed to read second sheet", http.StatusInternalServerError)
		return
	}
	rowsPersons, err := f.GetRows(sheets[2])
	if err != nil {
		http.Error(w, "Failed to read third sheet", http.StatusInternalServerError)
		return
	}
	sheetname := strings.ReplaceAll(sheets[1], ".", "/")
	sheetname = strings.ReplaceAll(sheetname, "-", "/")
	yearstr := strings.Split(sheetname, "/")[1]
	yearnum, err := strconv.Atoi(yearstr)
	if err != nil {
		log.Warn(r.Context(), "unable to parse year", zap.Error(err))
	}
	persons, err := application.ParseUsers(rowsPersons, r.Context())
	if err != nil {
		log.Error(r.Context(), "Error while parsing persons",
			zap.Error(err))
		return
	}
	electives, err := application.ParseElectives(rowsElectives, r.Context(), yearnum, persons)
	if err != nil {
		log.Error(r.Context(), "Error while parsing electives",
			zap.Error(err))
		return
	}

	courses, err := application.ParseCourses(rowsCourses, r.Context(), yearnum, persons)
	if err != nil {
		log.Error(r.Context(), "Error while parsing courses",
			zap.Error(err))
		return
	}
	allcourses := append(*electives, *courses...)
	resp := &Resp{
		Courses: allcourses,
		Persons: *persons,
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Error(r.Context(), "Failed to create JSON response", zap.Error(err))
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Error(r.Context(), "Failed to write response", zap.Error(err))
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
