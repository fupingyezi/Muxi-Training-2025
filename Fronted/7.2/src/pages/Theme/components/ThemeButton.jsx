import { useThemeContext } from '../index.jsx'

export function ThemeButton() {
  const { toggleTheme } = useThemeContext();
  
  return (
    <button className="theme-action-btn" onClick={toggleTheme}>
      ThemeButton切换主题
    </button>
  )
}