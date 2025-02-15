import { Box, Container } from "@gravity-ui/uikit";
import { NgwIdentify } from "@nextgis/ngw-kit";
import useMap from "features/hooks/useMap";
import { useEffect } from "react";

const Index = () => {
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

    const clean = () => {
      map.removeLayer("highlight");
    };

    const drawLayer = (identify: NgwIdentify) => {
			console.log(identify)
      map
        .fetchIdentifyGeoJson(identify)
        .then((geojson) => {
					console.log(geojson)
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

  return (
    <Container>
      <Box height={"100vh"} id="map"></Box>
    </Container>
  );
};

export default Index;
