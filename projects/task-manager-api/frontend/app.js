const API_URL = 'http://localhost:8080/tasks';

async function fetchTasks() {
    try {
        const response = await fetch(API_URL);
        const tasks = await response.json();

        const tasksDiv = document.getElementById('tasks');
        tasksDiv.innerHTML = '';

        tasks.forEach(task => {
            const taskDiv = document.createElement('div');
            taskDiv.className = 'task';

            taskDiv.innerHTML = `
                <h3>${task.title}</h3>
                <p><strong>Description:</strong> ${task.description}</p>
                <p><strong>Status:</strong> ${task.status}</p>
            `;

            tasksDiv.appendChild(taskDiv);
        });
    } catch (error) {
        console.error('Error fetching tasks:', error);
    }
}
