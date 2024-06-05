import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private baseUrl = 'http://localhost:8080/';

  private user: any;

  constructor(private http: HttpClient, private router: Router) { }

  register(data:any){
    return this.http.post<any>(`${this.baseUrl}users/register`,data)
  }
  login(data:any){
    return this.http.post<any>(`${this.baseUrl}users/login`,data)
  }
  signOut(){
    localStorage.clear();
    this.router.navigate(["login"]);
  }
  storeToken(tokenValue: any){
    this.user = tokenValue.user;
    localStorage.setItem("token", tokenValue.access_token);
    localStorage.setItem("id", this.user.id);
    localStorage.setItem("signed", "true");

    console.log(this.user);
  }
  getToken(){
    return localStorage.getItem("token");
  }
  isLoggedIn():boolean{
    return localStorage.getItem("signed") !== null;
  }
  isAdmin():boolean{
    return this.user.is_admin.Bool;
  }

  getUserId():number{
    const a = localStorage.getItem("id");

    if(a === null){
      return 0;
    }
    return +a;
  }
}
