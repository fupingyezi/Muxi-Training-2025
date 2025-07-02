import React, { useReducer, useEffect, useMemo, useRef } from 'react';
import { Link } from 'react-router-dom';
import "./index.css"

const TodoContext = React.createContext();

function todoReducer(state, action) {
  switch (action.type) {
    case 'ADD_TODO':
      return [...state, { id: Date.now(), text: action.text, completed: false }];
    case 'DELETE_TODO':
      return state.filter(todo => todo.id !== action.id);
    case 'TOGGLE_TODO':
      return state.map(todo =>
        todo.id === action.id ? { ...todo, completed: !todo.completed } : todo
      );
    default:
      return state;
  }
}

function TodoApp() {
  const [todos, dispatch] = useReducer(todoReducer, []);
  const inputRef = useRef(null);

  useEffect(() => {
    console.log('Todo List已加载');
  }, []);

  const handleAddTodo = () => {
    const text = inputRef.current.value;
    if (text.trim()) {
      dispatch({ type: 'ADD_TODO', text });
      inputRef.current.value = '';
    }
  };

  const handleDeleteTodo = (id) => {
    dispatch({ type: 'DELETE_TODO', id });
  };

  const handleToggleTodo = (id) => {
    dispatch({ type: 'TOGGLE_TODO', id });
  };

  const TodoList = useMemo(() => (
    <ul>
      {todos.map(todo => (
        <li key={todo.id}>
          <input
            type="checkbox"
            checked={todo.completed}
            onChange={() => handleToggleTodo(todo.id)}
          />
          {todo.text}
          <button onClick={() => handleDeleteTodo(todo.id)}>删除</button>
        </li>
      ))}
    </ul>
  ), [todos]);

  return (
    <TodoContext.Provider value={{ todos, dispatch }}>
      <div className="todo-container">
        <div className="todo-input-group">
          <input className="todo-input" ref={inputRef} type="text" placeholder="添加任务" />
          <button className="add-btn" onClick={handleAddTodo}>添加任务</button>
        </div>
        
        {TodoList}

        <div className="nav-group">
          <Link to="/" className="nav-btn">🔢 计数器答案</Link>
          <Link to="/theme" className="nav-btn">🎨 主题切换答案</Link>
        </div>
      </div>
    </TodoContext.Provider>
  );
}

export default TodoApp;