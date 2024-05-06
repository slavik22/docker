import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {

  signupForm!:FormGroup;
  constructor(private fb: FormBuilder, private authService: AuthService, private router: Router) { }

  ngOnInit(): void {
    this.signupForm = this.fb.group({
      name: ["",Validators.required],
      email: ["",Validators.required],
      password: ["",Validators.required],
    });
  }
  onSubmit(){
    if(this.signupForm.valid){
      this.authService.register(this.signupForm.value).subscribe({
        next: res => {
          this.signupForm.reset();
          this.router.navigate(['login']);
        },
        error: error => {
          alert(error.error.message)

        }
      })
    }
  }
}
