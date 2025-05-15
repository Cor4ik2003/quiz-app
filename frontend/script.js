const API_BASE = "http://localhost:8081"; // измени на свой URL

function login() {
  const email = document.getElementById("email").value;
  const password = document.getElementById("password").value;

  fetch(`${API_BASE}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password })
  })
    .then(res => {
      if (!res.ok) throw new Error("Неверный логин или пароль");
      return res.json();
    })
    .then(data => {
      localStorage.setItem("token", data.token);
      showQuizzes();
    })
    .catch(err => alert(err.message));
}

function logout() {
  localStorage.removeItem("token");
  document.getElementById("quiz-section").classList.add("d-none");
  document.getElementById("login-section").classList.remove("d-none");
}

function showQuizzes() {
  document.getElementById("login-section").classList.add("d-none");
  document.getElementById("quiz-section").classList.remove("d-none");

  fetch(`${API_BASE}/quizzes`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`
    }
  })
    .then(res => res.json())
    .then(data => {
      const container = document.getElementById("quizzes");
      container.innerHTML = "";

      data.forEach(quiz => {
        const col = document.createElement("div");
        col.className = "col-12 col-md-6";

        col.innerHTML = `
          <div class="card shadow-sm h-100">
            <div class="card-body">
              <h5 class="card-title">${quiz.title}</h5>
              <p class="card-text">${quiz.description || "Нет описания"}</p>
              <button class="btn btn-primary" onclick="startQuiz('${quiz.id}')">Пройти</button>
            </div>
          </div>
        `;
        container.appendChild(col);
      });
    })
    .catch(err => {
      alert("Ошибка при получении викторин");
      console.error(err);
    });
}
