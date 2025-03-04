import React, { useState, useEffect } from 'react';
import TaskForm from './components/TaskForm';
import TaskList from './components/TaskList';
import axios from 'axios';
import './App.css';

function App() {
  const [tasks, setTasks] = useState([]);
  const [currentTask, setCurrentTask] = useState(null);

  const fetchTasks = async () => {
    try {
      const response = await axios.get('http://localhost:8000/api/tasks');
      setTasks(response.data);
    } catch (err) {
      console.error('Error fetching tasks:', err);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  return (
    <div className="App">
      <h1>Task Manager</h1>
      <TaskForm currentTask={currentTask} setCurrentTask={setCurrentTask} refreshTasks={fetchTasks} />
      <TaskList tasks={tasks} setCurrentTask={setCurrentTask} refreshTasks={fetchTasks} />
    </div>
  );
}

export default App;
