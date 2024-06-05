import { Component } from '@angular/core';
import { TutorialService } from '../tutorial.service';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent {
  tutorials: any;
  userId: number = 1; // Replace with actual user ID logic

  constructor(private tutorialService: TutorialService, private userService: AuthService,  private router: Router) { }

  ngOnInit() {
    this.userId = this.userService.getUserId();
    this.tutorialService.getUserTutorials(this.userId).subscribe(data => {
      this.tutorials = data
    });
  }

  deleteTutorial(id: number) {
    this.tutorialService.deleteTutorial(id).subscribe(data => {
      window.location.reload();
    });
  }

  updateTutorial(id: number) {
    this.router.navigate(['/update', id]);
  }
}
