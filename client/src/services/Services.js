import api from "./api";

const Services = {
  Calculate: async (data) => {
    let response = {}
    await api
      .post(`/paint/calculate`, data)
      .then((res) => {
        response.data = res.data
      })
      .catch((error) => {
        response.error = error.response.data.error
      });
    return response;
  },
};

export default Services;
