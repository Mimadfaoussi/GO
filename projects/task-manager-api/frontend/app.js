const API_URL = 'http://localhost:8080/tasks';

async function fetchTasks() {
    try {
        const response = await fetch(API_URL);
        const tasks = await response.json();
        const tasksDiv = document.getElementById('tasks');
        tasksDiv.innerHTML = tasks.map(task => `<p>${task.title} - ${task.status}</p>`).join('');
    } catch (error) {
        console.error('Error fetching tasks:', error);
    }

}
