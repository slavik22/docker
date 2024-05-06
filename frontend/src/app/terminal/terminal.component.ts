import { AfterViewInit, Component, OnInit, ViewChild } from '@angular/core';
import { NgTerminal } from 'ng-terminal';
import { WebSocketService } from '../websocket.service';

@Component({
  selector: 'app-terminal',
  templateUrl: './terminal.component.html',
  styleUrls: ['./terminal.component.css']
})
export class TerminalComponent implements AfterViewInit, OnInit {
  @ViewChild('term', { static: false }) terminal!: NgTerminal;

  private wsSubscription! : any;
  private command: string = "";

  constructor(private wsService: WebSocketService) {}

  ngOnInit(): void {
    const wsMessages = this.wsService.connect('ws://localhost:8080/ws');
    this.wsSubscription = wsMessages.subscribe(
      message => {
        console.log('Received:',  message)
        this.terminal.write("\r\n");
        this.terminal.write(message);
        this.terminal.write("\r\n");
      } ,
      error => console.error(error),
      () => console.log('Completed!')
    );
  }

  ngAfterViewInit(){
    this.terminal.onData().subscribe((input) => {
      if (input === '\r') {
        this.wsService.send(this.command);
        this.command = "";
      } else if (input === '\u007f') {
          this.command = this.command.slice(0, -1);
          this.terminal.write('\b \b');
      } else if (input === '\u0003') {
          this.terminal.write('^C');
          this.terminal.write("\r\n");
      }else{
        this.terminal.write(input);
        this.command += input;
      }
    });
  }

  ngOnDestroy() {
    this.wsSubscription.unsubscribe();
  }

  // command(event: { key: string; domEvent: KeyboardEvent }): void {
  //   console.log(event.key)
  //   if (event.domEvent.key === 'Enter') {
      
  //     // this.terminal.write('\r\n');
  //   }
  // }
}
