import { Box, Container } from "@gravity-ui/uikit";
import { NgwIdentify } from "@nextgis/ngw-kit";
import useMap from "features/hooks/useMap";
import { useEffect } from "react";

const Map = () => {
  const map = useMap();

  useEffect(() => {
    const vectorLayerStyle = 8892;
    map.addNgwLayer({
      resource: vectorLayerStyle,
      fit: true,
      adapterOptions: {
        selectable: true,
      },
    });

    const drawLayer = (identify: NgwIdentify) => {
      console.log(identify);
      map
        .fetchIdentifyGeoJson(identify)
        .then((geojson) => {
          console.log(geojson);
        })
        .catch((e) => {
          if (e.name !== "AbortError") {
            throw e;
          }
        });
    };

    map.emitter.on("ngw:select", drawLayer);

    return () => {
      map.emitter.off("ngw:select", drawLayer);
    };
  }, []);

  return <Box height={"80vh"} id="map"></Box>;
};

export default Map;
