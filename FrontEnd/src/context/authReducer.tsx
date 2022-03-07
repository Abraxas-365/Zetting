import { User } from "../interfaces/appInterfaces";

export interface AuthState {
    status: 'checking' | 'autenticated' | 'not-autenticated';
    token: string | null;
    errorMessage: string | null;
    user: User | null;
    exists: boolean | null;
}

export type AuthAction =
    | { type: 'signUp', payload: { token: string, user: User } }
    | { type: 'signIn', payload: { token: string, user: User } }
    | { type: 'registerIn', payload: { exist: boolean, msg: string } }
    | { type: 'addError', payload: string }
    | { type: 'removeError' }
    | { type: 'notAutenticated' }
    | { type: 'logOut' }
    | { type: 'cleanEexists' }



export const authReducer = (state: AuthState, action: AuthAction): AuthState => {

    switch (action.type) {
        case 'cleanEexists':
            return {
                ...state,
                user: null,
                exists: null,
                token: null,
                errorMessage: null,
                status: 'not-autenticated'

            }
        case 'signIn':
        case 'signUp':
            return {
                ...state,
                errorMessage: null,
                status: 'autenticated',
                token: action.payload.token,
                user: action.payload.user
            }
        case 'notAutenticated':
        case 'logOut':
            return {
                ...state,
                status: 'not-autenticated',
                token: null,
                user: null
            }
        case "addError":
            return {
                ...state,
                user: null,
                status: 'not-autenticated',
                token: null,
                errorMessage: action.payload,

            }
        case "removeError":
            return {
                ...state,
                errorMessage: null,
            }
        default:
            return state;

    }

}
