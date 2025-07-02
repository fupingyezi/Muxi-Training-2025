import { useThemeContext } from '../index.jsx'

export function ThemeCard() {
  const { theme } = useThemeContext();
  
  return (
    <div className={`theme-card ${theme}`}>
      <h3>主题卡片</h3>
      <p>当前主题：{theme === 'light' ? '明亮' : '暗黑'}</p>
    </div>
  )
}