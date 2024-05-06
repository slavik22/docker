// src/app/websocket.service.ts
import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class WebSocketService {
  private socket!: WebSocket;
  private messageSubject: Subject<string> = new Subject();

  constructor() {}

  public connect(url: string): Subject<string> {
      if (!this.socket || this.socket.readyState === WebSocket.CLOSED) {
          this.socket = new WebSocket(url);

          this.socket.onmessage = (event) => {
              this.messageSubject.next(event.data);
          };

          this.socket.onerror = (event) => {
              console.error('WebSocket error:', event);
          };

          this.socket.onclose = (event) => {
              console.log('WebSocket connection closed:', event);
          };
      }

      return this.messageSubject;
  }

  public send(message: string): void {
    console.log("Sending: " + message)
      if (this.socket.readyState === WebSocket.OPEN) {
          this.socket.send(message);
      } else {
          console.error('WebSocket is not open.');
      }
  }
}
