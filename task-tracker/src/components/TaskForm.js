import React, { useState, useEffect } from 'react';
import axios from 'axios';
import "./TaskForm.css";

const TaskForm = ({ currentTask, setCurrentTask, refreshTasks }) => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [dueDate, setDueDate] = useState('');
  const [status, setStatus] = useState('Pending');

  useEffect(() => {
    if (currentTask) {
      setTitle(currentTask.title);
      setDescription(currentTask.description);
      setDueDate(currentTask.dueDate);
      setStatus(currentTask.status);
    }
  }, [currentTask]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const taskData = { title, description, dueDate, status };

    try {
      if (currentTask) {
        await axios.put(`http://localhost:8000/api/tasks/${currentTask._id}`, taskData);
      } else {
        await axios.post('http://localhost:8000/api/tasks', taskData);
      }
      refreshTasks();
      setTitle('');
      setDescription('');
      setDueDate('');
      setStatus('Pending');
      setCurrentTask(null);
    } catch (err) {
      console.error('Error saving task:', err);
      alert(`Failed to save task: ${err.response ? err.response.data : err.message}`);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Title"
        required
      />
      <input
        type="text"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
        placeholder="Description"
        required
      />
      <input
        type="date"
        value={dueDate}
        onChange={(e) => setDueDate(e.target.value)}
        required
      />
      <select value={status} onChange={(e) => setStatus(e.target.value)}>
        <option value="Pending">Pending</option>
        <option value="In Progress">In Progress</option>
        <option value="Completed">Completed</option>
      </select>
      <button type="submit">{currentTask ? 'Update' : 'Add'}</button>
    </form>
  );
};

export default TaskForm;
