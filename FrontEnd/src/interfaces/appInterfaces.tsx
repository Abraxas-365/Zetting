
export interface LoginResponse {
    user: User;
    token: string;
}

export interface User {

    id?: string;
    first_name?: string;
    last_name?: string;
    email?: string;
    phone?: string;
    country?: string;
    professions?: Array<string>;
    identifierDocument: string;
    password: string;
    verified: boolean;
    img?: string;


}

export interface Project {
    name: string;
    description?: string;
    calendar?: string;
    collaboration?: string;
    color?: string;
    empleados?: Array<string>;


}

export interface Features {
    id?: string;
    userId?: string;
    height?: number;
    body?: string;
    Skin?: string;
    hair_type?: Array<string>;
    hair_color?: string;
    eye_color?: string;


}
