import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private baseUrl = 'http://localhost:8080/';

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
  storeToken(tokenValue: string){
    localStorage.setItem("token", tokenValue);
  }
  getToken(){
    return localStorage.getItem("token");
  }
  isLoggedIn():boolean{
    return !!localStorage.getItem("token");
  }
}
