const routes: Routes = [
    { path: 'login', component: LoginComponent },
    { path: 'register', component: RegisterComponent },
    { path: 'student', component: StudentComponent },
    { path: 'teacher', component: TeacherComponent },
    { path: '**', redirectTo: 'login' }
  ];
  