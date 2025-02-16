// Map.tsx
import { useState } from "react";
import useMap, { OnSelectFeatureCallback } from "features/hooks/useMap";
import { Modal, Text, Button } from "@gravity-ui/uikit";

const Map = () => {
  const [modalOpen, setModalOpen] = useState(false);
  const [soldierData, setSoldierData] = useState<any>(null);

  // Колбэк, вызываемый при клике по объекту на карте
  const handleSelectFeature: OnSelectFeatureCallback = (featureData, coordinate) => {
    console.log(coordinate);
    setSoldierData(featureData.feature);
    setModalOpen(true);
  };

  // Инициализируем карту и передаём колбэк для выбора объекта
  useMap(handleSelectFeature);

  return (
    <div>
      {/* Контейнер карты */}
      <div id="map" style={{ height: "80vh" }}></div>

      {/* Модальное окно от Gravity UI */}
      <Modal open={modalOpen} onClose={() => setModalOpen(false)}>
        {soldierData && (
          <div style={{ padding: "20px", display: "flex", flexDirection: "column", gap: "10px" }}>
            <Text variant="display-3">
              {soldierData.fio}
            </Text>
            <Text variant="body-3">
              <strong>Район:</strong> {soldierData.n_raion}
            </Text>
            <Text variant="body-3">
              <strong>Годы жизни:</strong> {soldierData.years}
            </Text>
            <Text variant="body-3">
              <strong>Информация:</strong> {soldierData.info}
            </Text>
            <Text variant="body-3">
              <strong>Контракт:</strong> {soldierData.kontrakt}
            </Text>
            <Text variant="body-3">
              <strong>Награды:</strong> {soldierData.nagrads}
            </Text>
            <Button
              view="action"
              size="l"
              onClick={() => window.location.assign(`/person/${soldierData.id}`)}
            >
              Перейти к солдату
            </Button>
          </div>
        )}
      </Modal>
    </div>
  );
};

export default Map;
