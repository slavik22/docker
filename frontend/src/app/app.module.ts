import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClient, HttpClientModule } from '@angular/common/http'
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgTerminalModule } from 'ng-terminal';
import { TerminalComponent } from './terminal/terminal.component';
import { TutorialComponent } from './tutorial/tutorial.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MainComponent } from './main/main.component';
import { TutorialListComponent } from './tutorial-list/tutorial-list.component';
import { HeaderComponent } from './header/header.component';
import { AddTutorialComponent } from './add-tutorial/add-tutorial.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { UpdateTutorialComponent } from './update-tutorial/update-tutorial.component';

@NgModule({
  declarations: [
    AppComponent,
    TerminalComponent,
    TutorialComponent,
    LoginComponent,
    RegisterComponent,
    MainComponent,
    TutorialListComponent,
    HeaderComponent,
    AddTutorialComponent,
    DashboardComponent,
    UpdateTutorialComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    NgTerminalModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
