import React, { useState, useEffect, useCallback } from "react";
import { Link } from "react-router-dom";
import "./index.css";

// 自定义Hook，useCounter
function useCounter(initialValue = 0) {
  const [count, setCount] = useState(initialValue);

  const increment = useCallback(() => {
    setCount((prevCount) => prevCount + 1);
  }, []);

  const decrement = useCallback(() => {
    setCount((prevCount) => prevCount - 1);
  }, []);

  const reset = useCallback(() => {
    setCount(initialValue);
  }, [initialValue]);

  useEffect(() => {
    console.log(`计数值已改变：${count}`);
  }, [count]);

  return { count, increment, decrement, reset };
}

function CounterApp() {
  const { count, increment, decrement, reset } = useCounter(0);

  return (
    <div className="counter-container">
      <h1 className="counter-title">计数器答案</h1>
      <div className="counter-display">{count}</div>
      <div className="button-group">
        <button className="increment-btn" onClick={increment}>
          +
        </button>
        <button className="decrement-btn" onClick={decrement}>
          -
        </button>
        <button className="reset-btn" onClick={reset}>
          ↻ 重置
        </button>
        <Link to="/theme" className="nav-btn">
          🎨 主题切换答案
        </Link>
        <Link to="/todo" className="nav-btn">
          📝 TodoList答案
        </Link>
      </div>
    </div>
  );
}

export default CounterApp;
