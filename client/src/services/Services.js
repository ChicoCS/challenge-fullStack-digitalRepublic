import api from "./api";

//FIX ROUTE!
const Services = {
  Calculate: async (data) => {
    let response = {}
    await api
      .post(`/paint`, data)
      .then((res) => {
        response.data = res.data
      })
      .catch(({ res }) => {
        response.error = res.data
      });
    return response;
  },
};

export default Services;
