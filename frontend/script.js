const API_URL = "http://localhost:8082"; // измени порт, если у тебя другой

async function createQuiz() {
  console.log("Кнопка нажата");

  const title = document.getElementById("title").value;
  if (!title) {
    alert("Введите название квиза");
    return;
  }

  const response = await fetch(`${API_URL}/quizzes`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ title: title })
  });

  if (response.ok) {
    alert("Квиз создан!");
    fetchQuizzes(); // обновим список
  } else {
    const text = await response.text();
    alert("Ошибка при создании квиза: " + text);
  }
}

async function fetchQuizzes() {
  const res = await fetch(`${API_URL}/quizzes`);
  const quizzes = await res.json();
  const list = document.getElementById("quiz-list");
  list.innerHTML = "";
  quizzes.forEach(q => {
    const li = document.createElement("li");
    li.innerHTML = `
      <strong>${q.title}</strong>
      <button onclick="deleteQuiz('${q.id}')">Delete</button>
    `;
    list.appendChild(li);
  });
}

async function deleteQuiz(id) {
  const res = await fetch(`${API_URL}/quizzes/${id}`, {
    method: "DELETE"
  });

  if (res.ok) {
    alert("Удалено");
    fetchQuizzes();
  } else {
    alert("Не удалось удалить квиз");
  }
}

fetchQuizzes();
