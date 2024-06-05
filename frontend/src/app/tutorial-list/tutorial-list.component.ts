import { Component } from '@angular/core';
import { TutorialService } from '../tutorial.service';

@Component({
  selector: 'app-tutorial-list',
  templateUrl: './tutorial-list.component.html',
  styleUrls: ['./tutorial-list.component.css']
})
export class TutorialListComponent {
  tutorials: any;

  constructor(private tutorialService: TutorialService) { }

  ngOnInit() {
    this.tutorialService.getTutorials().subscribe(data => {
      this.tutorials = data;
    });
  }
}
