import { Box } from "@gravity-ui/uikit";
import useMap from "features/hooks/useMap";

const Map = () => {
  useMap();

  return <Box height={"80vh"} id="map"></Box>;
};

export default Map;
