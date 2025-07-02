import React, {
  useState,
  useEffect,
  useCallback,
  createContext,
  useContext,
} from "react";
import { Link } from "react-router-dom";
import { ThemeCard } from "./components/ThemeCard";
import { ThemeButton } from "./components/ThemeButton";
import "./index.css";

const ThemeContext = createContext();

export function useThemeContext() {
  const context = useContext(ThemeContext);
  if (!context) {
    throw new Error("useThemeContext must be used within a ThemeProvider");
  }
  return context;
}

function useTheme(initialTheme = "light") {
  const [theme, setTheme] = useState(initialTheme);

  const toggleTheme = useCallback(() => {
    setTheme((prevTheme) => (prevTheme === "light" ? "dark" : "light"));
  }, []);

  useEffect(() => {
    document.body.style.backgroundColor = theme === "light" ? "#fff" : "#000";
    document.body.style.color = theme === "light" ? "#000" : "#fff";
  }, [theme]);

  return { theme, toggleTheme };
}

function ThemeApp() {
  const { theme, toggleTheme } = useTheme("light");

  return (
    <ThemeContext.Provider value={{ theme, toggleTheme }}>
      <div className={`theme-container ${theme}`}>
        <ThemeCard />
        <ThemeButton />
        <h1 className="theme-title">
          {theme === "light" ? "🌞 明亮模式" : "🌙 暗黑模式"}
        </h1>
        <button className="theme-toggle-btn" onClick={toggleTheme}>
          🎨 ThemeApp切换主题
        </button>
        <Link to="/" className="nav-btn">
          📝 Counter答案
        </Link>
        <Link to="/todo" className="nav-btn">
          📝 TodoList答案
        </Link>
      </div>
    </ThemeContext.Provider>
  );
}

export default ThemeApp;
