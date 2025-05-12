@Component({
  selector: 'app-register',
  template: `
    <h2>Регистрация</h2>
    <form (submit)="onRegister($event)">
      <input [(ngModel)]="email" name="email" placeholder="Email">
      <input [(ngModel)]="password" name="password" placeholder="Пароль" type="password">
      <select [(ngModel)]="role" name="role">
        <option value="student">Студент</option>
        <option value="teacher">Преподаватель</option>
      </select>
      <button type="submit">Зарегистрироваться</button>
    </form>
  `
})
export class RegisterComponent {
  email = '';
  password = '';
  role = 'student';

  constructor(private auth: AuthService, private router: Router) {}

  onRegister(e: Event) {
    e.preventDefault();
    this.auth.register(this.email, this.password, this.role).subscribe({
      next: () => this.router.navigate(['/login']),
      error: err => alert('Ошибка регистрации')
    });
  }
}
