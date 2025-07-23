export type InstituteType = {
  label: string,
  value: number
}

export const instituteList: InstituteType[] = [
  { value: 1, label: 'Институт анализа данных и Искусственного Интеллекта' },
  { value: 2, label: 'Институт разработки ПО и программной инженерии' },
  { value: 3, label: 'Институт робототехники и компьютерного зрения' },
  { value: 4, label: 'Институт информационной безопасности' },
  { value: 5, label: 'Институт гуманитарных и социальных наук' }
]

export type RoleType = {
  value: number,
  label: string
}

export const roleList: RoleType[] = [
  { value: 1, label: 'Professor' },
  { value: 2, label: 'Docent' },
  { value: 3, label: 'Senior Instructor' },
  { value: 4, label: 'Instructor' },
  { value: 5, label: 'TA' },
  { value: 6, label: 'TA intern' },
  { value: 7, label: 'Visiting' }
]
