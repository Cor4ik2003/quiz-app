import { Component } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  template: `
    <h2>Вход</h2>
    <form (submit)="onLogin($event)">
      <input [(ngModel)]="email" name="email" placeholder="Email">
      <input [(ngModel)]="password" name="password" placeholder="Пароль" type="password">
      <button type="submit">Войти</button>
    </form>
  `
})
export class LoginComponent {
  email = '';
  password = '';

  constructor(private auth: AuthService, private router: Router) {}

  onLogin(e: Event) {
    e.preventDefault();
    this.auth.login(this.email, this.password).subscribe({
      next: (res: any) => {
        this.auth.setToken(res.token);
        const role = this.auth.getUserRole();
        if (role === 'student') this.router.navigate(['/student']);
        else if (role === 'teacher') this.router.navigate(['/teacher']);
      },
      error: err => alert('Ошибка входа')
    });
  }
}
