import { useState } from "react";
import "./styles/App.css";
import Game from "./pages/Game/Game";
import Menu from "./pages/Menu/Menu";

export default function App() {
  const [gameMode, setGameMode] = useState<"MENU" | "PVP" | "PVE">("MENU");

  const handleSelectMode = (mode: "PVP" | "PVE") => {
    setGameMode(mode);
  };

  return (
    <>
      {gameMode === "MENU" ? (
        <Menu onSelectMode={handleSelectMode} />
      ) : (
        <Game />
      )}
    </>
  );
}