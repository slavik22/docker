import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private auth: AuthService, private router: Router){}

  canActivate():boolean {
      if(this.auth.isLoggedIn()){
        return true;
      }
      else{
        console.log("Please login first");
        this.router.navigate(['login']);
        return  false;
  
      }
  }
  
}
