export interface ILoginForm {
  email: string
  password: string
}

export interface ILoginFormValidate {
  email: Object
  password: Object
}

export const LoginFormValidate: ILoginFormValidate = {
  email: {
    required: true,
    email: true,
  },
  password: {
    required: true,
  },
}
