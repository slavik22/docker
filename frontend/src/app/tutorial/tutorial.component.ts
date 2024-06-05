import { Component, Input, OnInit } from '@angular/core';
import { TutorialService } from '../tutorial.service';

@Component({
  selector: 'app-tutorial',
  templateUrl: './tutorial.component.html',
  styleUrls: ['./tutorial.component.css']
})
export class TutorialComponent implements OnInit{
  @Input('id') id = 0;
  tutorial: any;

  constructor(private tutorialService: TutorialService) { }

  ngOnInit() {
    console.log ("before first for loop this is maxrowscell " + this.id);
    this.tutorialService.getTutorial(this.id).subscribe(data => {
      this.tutorial = data;
    });
  }

}
