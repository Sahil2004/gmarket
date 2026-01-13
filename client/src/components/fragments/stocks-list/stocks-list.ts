import { Component, Input } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatListModule } from '@angular/material/list';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'stocks-list',
  templateUrl: 'stocks-list.html',
  imports: [MatListModule, RouterLink, MatButtonModule],
})
export class StocksList {
  @Input({ required: true }) stocks: string[] = [];
  @Input() displayActions: boolean = true;
  @Input() buyHandler: (stock: string) => void = () => {};
  @Input() sellHandler: (stock: string) => void = () => {};

  activated(stock: string): boolean {
    return false;
  }
}
