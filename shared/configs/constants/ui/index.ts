export type InstituteType = {
  name: string,
  id: number
}

export const instituteList: InstituteType[] = [
  { id: 1, name: 'Институт анализа данных и Искусственного Интеллекта' },
  { id: 2, name: 'Институт разработки ПО и программной инженерии' },
  { id: 3, name: 'Институт робототехники и компьютерного зрения' },
  { id: 4, name: 'Институт информационной безопасности' },
  { id: 5, name: 'Институт гуманитарных и социальных наук' }
]

export type RoleType = {
  id: number,
  name: string
}

export const roleList: RoleType[] = [
  { id: 1, name: 'Professor' },
  { id: 2, name: 'Docent' },
  { id: 3, name: 'Senior Instructor' },
  { id: 4, name: 'Instructor' },
  { id: 5, name: 'TA' },
  { id: 6, name: 'TA intern' },
  { id: 7, name: 'Visiting' }
]
