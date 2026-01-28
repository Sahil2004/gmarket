import { Component, inject, Input } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog } from '@angular/material/dialog';
import { ImageUploadDialog } from '../..';

@Component({
  selector: 'input-image',
  templateUrl: 'input-image.html',
  imports: [MatButtonModule],
})
export class InputImage {
  @Input() changeImage: boolean = true;
  @Input() imageUri: string | null = null;
  @Input() defaultText: string = 'No Image';
  @Input({ required: true }) title!: string;
  @Input({ required: true }) uploadHandler!: (newImage: string) => boolean;

  _dialog = inject(MatDialog);

  openImageUploadDialog() {
    console.log(this.imageUri);
    this._dialog.open(ImageUploadDialog, {
      data: {
        title: this.title,
        imageUri: this.imageUri,
        defaultText: this.defaultText,
        upload: (newImage: string) => this.uploadHandler(newImage),
      },
    });
  }
}
