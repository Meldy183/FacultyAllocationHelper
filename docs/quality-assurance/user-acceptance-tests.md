# ✅ Acceptance Tests

This section includes user acceptance tests written in the GIVEN/WHEN/THEN format. Each scenario is linked to a corresponding issue (if applicable) and reflects whether it was accepted or is still in discussion.

---

## 📘 Simplify the Course's Card — Add Track Information  
**🔗 Issue**: #103 — ✅ *Accepted by the customer*

**Acceptance Criteria**  
- **GIVEN** I am an authorized user with access to course information  
- **WHEN** I am viewing a course’s card  
- **THEN** I can see which program (track) the course belongs to  

---

## 💬 Use Pop-up Forms Instead of Redirections  
**🔗 Issue**: #104 — ✅ *Accepted by the customer*

**Acceptance Criteria**  
- **GIVEN** I am an authorized user able to manage TAs and courses  
- **WHEN** I update the information about a course or a TA  
- **THEN** The form opens in a pop-up window and I am **not redirected** to another page  

---

## 🔁 Refactor Workload from Hours to Percent  
**🔗 Issue**: #117 — ✅ *Accepted by the customer*

**Acceptance Criteria**  
- **GIVEN** I am an authorized user able to update faculty members  
- **WHEN** I update the workload of a faculty member  
- **THEN** The workload is shown and stored in **percent (%)** instead of **hours**  

---

## 🧠 Faculty–Course Accordance Prediction  
**📝 Status**: Discussed during the meeting — ❌ *Not planned yet*

**Acceptance Criteria**  
- **GIVEN** I am an authorized user able to manage TAs and courses  
- **WHEN** I am selecting a faculty member to assign to a course  
- **THEN** The list of faculties is **ordered by accordance**, calculated using a special formula  

---

## 🕒 "Time Machine" Feature (Read-only Past Year View)  
**📝 Status**: Discussed during the meeting — 🟡 *In progress*

**Acceptance Criteria**  
- **GIVEN** I am an authorized user able to manage TAs and courses  
- **WHEN** I am viewing data for the current academic year  
- **THEN** I can switch to a **previous year’s data** and **view (but not edit)** the same information  

---

> 💡 These acceptance tests reflect functional expectations directly aligned with stakeholder discussions and feedback. Any "not planned yet" items may be prioritized in future roadmap iterations.
