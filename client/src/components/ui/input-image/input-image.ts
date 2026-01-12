import { Component, Input } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'input-image',
  templateUrl: 'input-image.html',
  imports: [MatButtonModule],
})
export class InputImage {
  @Input() changeImage: boolean = true;
  @Input() imageUri: string | null = null;
  @Input() defaultText: string = 'No Image';
}
