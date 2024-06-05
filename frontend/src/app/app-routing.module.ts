import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './auth.guard';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { MainComponent } from './main/main.component';
import { TutorialListComponent } from './tutorial-list/tutorial-list.component';
import { AddTutorialComponent } from './add-tutorial/add-tutorial.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { UpdateTutorialComponent } from './update-tutorial/update-tutorial.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: '', component: TutorialListComponent, canActivate:[AuthGuard] },
  { path: 'tutorial/:id', component: MainComponent },
  { path: 'tutorial/:id', component: MainComponent },
  { path: 'update/:id', component: UpdateTutorialComponent },
  { path: 'dashboard', component: DashboardComponent },
  { path: 'add', component: AddTutorialComponent },
  { path: 'register', component: RegisterComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
