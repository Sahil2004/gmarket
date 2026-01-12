import { Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';

@Component({
  selector: 'confirmation-dialog',
  templateUrl: 'confirmation-dialog.html',
  imports: [MatDialogModule, MatButtonModule],
})
export class ConfirmationDialog {
  data = inject(MAT_DIALOG_DATA);
}
