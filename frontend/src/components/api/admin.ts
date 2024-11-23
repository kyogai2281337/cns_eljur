import axios from "axios";

export const backendURL = process.env.VUE_APP_BACKEND_URL || "/api";
const baseUrl = backendURL + "/auth";

interface getTablesStructure {
  tables: string[];
}

const getTables = async (): Promise<{
  data?: getTablesStructure;
  error: boolean;
  errorMsg?: string;
}> => {
  try {
    const response = await axios.get<getTablesStructure>(
      baseUrl + "/private/gettables"
    );
    return {
      data: response.data,
      error: false,
    };
  } catch (error) {
    if (axios.isAxiosError(error)) {
      return {
        error: true,
        errorMsg: error.response?.data?.message || "Failed to fetch tables",
      };
    }
    return {
      error: true,
      errorMsg: "Unexpected error occurred",
    };
  }
};

interface getListObjs {
  id: number;
  email: string;
}
interface getListStructure {
  table: getListObjs[];
}

const getList = async (
  tablesname: string,
  limit: number,
  page: number
): Promise<{ data?: any; error: boolean; errorMsg?: string }> => {
  try {
    const response = await axios.post<getListStructure>(
      baseUrl + "/private/getlist",
      {
        tablename: tablesname,
        limit: limit,
        page: page,
      }
    );
    return {
      data: response.data,
      error: false,
    };
  } catch (error) {
    if (axios.isAxiosError(error)) {
      return {
        error: true,
        errorMsg: error.response?.data?.message || "Failed to fetch list",
      };
    }
    return {
      error: true,
      errorMsg: "Unexpected error occurred",
    };
  }
};

const getObj = async (
  tablename: string,
  id: number
): Promise<{ data?: any; error: boolean; errorMsg?: string }> => {
  try {
    const response = await axios.post<any>(baseUrl + "/private/getobj", {
      tablename: tablename,
      id: id,
    });
    return {
      data: response.data,
      error: false,
    };
  } catch (error) {
    if (axios.isAxiosError(error)) {
      return {
        error: true,
        errorMsg: error.response?.data?.message || "Failed to fetch obj",
      };
    }
    return {
      error: true,
      errorMsg: "Unexpected error occurred",
    };
  }
};

const setObj = async (
  tablename: string,
  table: any
): Promise<{ data?: any; error: boolean; errorMsg?: string }> => {
  try {
    const response = await axios.post<any>(baseUrl + "/private/setobj", {
      tablename: tablename,
      table: table,
    });
    return {
      data: response.data,
      error: false,
    };
  } catch (error) {
    if (axios.isAxiosError(error)) {
      return {
        error: true,
        errorMsg: error.response?.data?.message || "Failed to set obj",
      };
    }
    return {
      error: true,
      errorMsg: "Unexpected error occurred",
    };
  }
};

export default { getTables, getList, getObj, setObj };
