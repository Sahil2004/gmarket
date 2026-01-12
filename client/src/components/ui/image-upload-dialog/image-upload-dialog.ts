import { Component, inject, signal } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'image-upload-dialog',
  templateUrl: 'image-upload-dialog.html',
  imports: [MatDialogModule, MatButtonModule],
})
export class ImageUploadDialog {
  data = inject(MAT_DIALOG_DATA);
  private _snackBar = inject(MatSnackBar);
  _imageUri = signal<string | null>(this.data.imageUri);
  onFileSelected(event: Event) {
    const input = event.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      const file = input.files[0];
      const reader = new FileReader();
      reader.onload = (e) => {
        this._imageUri.set(e.target?.result as string);
      };
      reader.readAsDataURL(file);
    }
  }
  upload() {
    if (!this._imageUri()) {
      this._snackBar.open('No image selected to upload', 'Close', { duration: 3000 });
    }
    this.data.upload(this._imageUri());
  }
}
