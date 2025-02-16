// useMap.tsx
import NgwMap from "@nextgis/ngw-ol";
import { useRef, useEffect } from "react";

export type OnSelectFeatureCallback = (
  featureData: { soldierId: string | number; feature: any; pixel: [number, number] },
  coordinate: [number, number]
) => void;

const useMap = (onSelectFeature?: OnSelectFeatureCallback) => {
  const mapRef = useRef<NgwMap | null>(null);

  useEffect(() => {
    if (!mapRef.current) {
      const options = {
        baseUrl: "/api-map",
        target: "map",
        auth: {
          login: "hackathon_19",
          password: "hackathon_19_25",
        },
        adapterOptions: {
          selectable: true,
        },
      };

      const newMap = new NgwMap(options);
      mapRef.current = newMap;

      const vectorLayerStyle = 8892;

      newMap.addNgwLayer({
        resource: vectorLayerStyle,
        fit: true,
        adapterOptions: {
          selectable: true,
        },
      });

      newMap
        .onLoad()
        .then(() => {
          console.log("Карта загружена");
        })
        .catch((e) => {
          console.error("Ошибка при загрузке карты:", e);
        });

      // При выборе (клике по объекту) получаем данные объекта
      newMap.emitter.on("ngw:select", async (e: any) => {
        if (e) {
          const items = e.getIdentifyItems();
          if (items && items.length > 0) {
            try {
              for (const item of items) {
                // Фильтруем объекты по родительскому слою
                if (item.parent !== "Хакатон-19") {
                  continue;
                }
                console.log("Полученные поля:", item.fields);
                const data = item.fields;
                const soldierId = data.id;
                if (onSelectFeature) {
                  // Передаём данные в колбэк.
                  // Здесь координаты и пиксели не используются для позиционирования модального окна, поэтому передаём [0, 0].
                  onSelectFeature({ soldierId, feature: data, pixel: [0, 0] }, [0, 0]);
                }
                break; // Если нужно обработать только первый подходящий объект
              }
            } catch (error) {
              console.error("Ошибка обработки выбранного объекта:", error);
            }
          }
        }
      });
    }
  }, [onSelectFeature]);

  return mapRef.current;
};

export default useMap;
