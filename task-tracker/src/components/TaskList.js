import React from 'react';
import axios from 'axios';
import "./TaskList.css";

const TaskList = ({ tasks, setCurrentTask, refreshTasks }) => {
  const handleDelete = async (id) => {
    try {
      await axios.delete(`http://localhost:8000/api/tasks/${id}`);
      refreshTasks();
    } catch (err) {
      console.error('Error deleting task:', err);
    }
  };

  return (
    <div style={{ padding: 0, margin: 0 }}>
      <h2>Task List</h2>
      {tasks && tasks.length > 0 ? (
        <ul>
          {tasks.map((task) => (
            <li key={task._id}>
              <span>{task.title} - {task.status}</span>
              <span style={{ display: "flex" }}>
                <button onClick={() => setCurrentTask(task)}>Edit</button>
                <button onClick={() => handleDelete(task._id)}>Delete</button>
              </span>
            </li>
          ))}
        </ul>
      ) : (
        <p>No Tasks are available.</p>
      )
      }
    </div >
  );
};

export default TaskList;
