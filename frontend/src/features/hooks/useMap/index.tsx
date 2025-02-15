import NgwMap from "@nextgis/ngw-ol";
const useMap = () => {
  const options = {
    baseUrl: "/api",
    target: "map",
    auth: {
      login: "hackathon_19",
      password: "hackathon_19_25",
    },
    adapterOptions: {
      selectable: true,
    }
  };
  const map = new NgwMap(options);

  return map;
};

export default useMap;
