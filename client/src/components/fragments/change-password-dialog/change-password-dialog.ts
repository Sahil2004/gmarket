import { Component, inject } from '@angular/core';
import {
  FormBuilder,
  ReactiveFormsModule,
  type ValidationErrors,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { matchFieldValidator, unmatchFieldValidator } from '../../../directives';

@Component({
  selector: 'change-password-dialog',
  templateUrl: 'change-password-dialog.html',
  imports: [
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    ReactiveFormsModule,
  ],
})
export class ChangePasswordDialog {
  data = inject(MAT_DIALOG_DATA);
  fb = inject(FormBuilder);

  changePasswordForm = this.fb.group(
    {
      oldPassword: ['', [Validators.required]],
      newPassword: [
        '',
        [
          Validators.required,
          Validators.minLength(6),
          Validators.pattern('^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{6,}$'),
        ],
      ],
      confirmNewPassword: ['', [Validators.required]],
    },
    {
      validators: [
        matchFieldValidator('confirmNewPassword', 'newPassword'),
        unmatchFieldValidator('newPassword', 'oldPassword'),
      ],
    }
  );

  get oldPassword() {
    return this.changePasswordForm.get('oldPassword');
  }

  get newPassword() {
    return this.changePasswordForm.get('newPassword');
  }

  get confirmNewPassword() {
    return this.changePasswordForm.get('confirmNewPassword');
  }

  getErrorMessage(controlName: string, errors: ValidationErrors | null | undefined): string {
    if (errors) {
      switch (Object.keys(errors)[0]) {
        case 'required':
          return `${controlName} is required`;
        case 'email':
          return `Not a valid email`;
        case 'minlength':
          return `${controlName} must be at least 6 characters long`;
        case 'pattern':
          return `${controlName} must contain at least one letter and one number`;
      }
    }
    if (this.changePasswordForm.errors) {
      switch (Object.keys(this.changePasswordForm.errors)[0]) {
        case 'fieldsMismatch':
          return `Passwords do not match`;
        case 'fieldsMatch':
          return `New password must be different from old password`;
      }
    }
    return '';
  }

  changePasswordHandler() {
    if (this.changePasswordForm.invalid) return;
    return this.data.onSubmit(this.oldPassword?.value, this.newPassword?.value);
  }
}
