import axios, {AxiosResponse} from 'axios';

interface GetTablesResponse {
    tables: string[];
}

interface GetListRequest {
    tablename: string;
    limit: number;
    page: number;
}

interface GetListItemResponse {
    id: number;
    email: string;
}

interface GetListResponse {
    table: GetListItemResponse[];
}

interface GetObjRequest {
    tablename: string;
    id: number;
}

interface GetObjRoleResponse {
    id: number;
    name: string;
}

interface GetObjResponse {
    id: number;
    email: string;
    first_name: string;
    last_name: string;
    role: GetObjRoleResponse;
    deleted: boolean;
}

interface SetObjRequest {
    tablename: string;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    table: any;
}

interface SetObjResponse {
    ID: number;
    Email: string;
    FirstName: string;
    LastName: string;
    Role: GetObjRoleResponse;
    IsActive: boolean;
}

interface CreateRequest {
    tablename: string;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: any;
}

export const getTables = async (): Promise<AxiosResponse<GetTablesResponse>> => {
    return await axios.get('/api/admin/private/gettables');
};

export const getList = async (
    request: GetListRequest
): Promise<AxiosResponse<GetListResponse>> => {
    return await axios.post('/api/admin/private/getlist', request);
};

export const getObj = async (
    request: GetObjRequest
): Promise<AxiosResponse<GetObjResponse>> => {
    return await axios.post('/api/admin/private/getobj', request);
};

export const setObj = async (
    request: SetObjRequest
): Promise<AxiosResponse<SetObjResponse>> => {
    return await axios.post('/api/admin/private/setobj', request);
};

export const create = async (
    request: CreateRequest
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
): Promise<AxiosResponse<any>> => {
    return await axios.post('/api/admin/private/create', request);
};
