import { Component } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

@Component({
  selector: 'change-password-dialog',
  templateUrl: 'change-password-dialog.html',
  imports: [MatFormFieldModule, MatInputModule, MatButtonModule, ReactiveFormsModule],
})
export class ChangePasswordDialog {}
