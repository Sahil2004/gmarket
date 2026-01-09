import { Component, inject } from '@angular/core';
import { MatGridListModule } from '@angular/material/grid-list';
import { InputImage } from '../../components';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'profile',
  templateUrl: 'profile.html',
  imports: [MatGridListModule, InputImage],
})
export class Profile {
  readonly userService = inject(UserService);

  get currentUser() {
    return this.userService.currentUser;
  }

  get nameInitials(): string {
    if (!this.currentUser) return '';
    const nameArr = this.currentUser.name.split(' ');
    const initials = nameArr.map((n) => n.charAt(0).toUpperCase());
    console.log(initials);
    return initials.join('');
  }
}
