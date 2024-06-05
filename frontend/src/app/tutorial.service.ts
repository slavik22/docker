import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})

export class TutorialService {
  private baseUrl = 'http://localhost:8080/tutorials';

  constructor(private http: HttpClient) { }

  getTutorials(): Observable<any> {
    return this.http.get(this.baseUrl);
  }

  getTutorial(id: number): Observable<any> {
    return this.http.get(`${this.baseUrl}/${id}`);
  }

  createTutorial(tutorial: Object): Observable<Object> {
    return this.http.post(this.baseUrl, tutorial);
  }

  updateTutorial(id: number, value: any): Observable<Object> {
    return this.http.put(`${this.baseUrl}/${id}`, value);
  }

  getUserTutorials(id: number): Observable<any> {
    return this.http.get(`${this.baseUrl}/user/${id}`);
  }
  deleteTutorial(id: number): Observable<Object> {
    return this.http.delete(`${this.baseUrl}/${id}`);
  }
}
