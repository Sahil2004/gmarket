import { Component, Input } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'stocks-list',
  templateUrl: 'stocks-list.html',
  imports: [MatListModule, RouterLink, MatButtonModule, MatIconModule],
})
export class StocksList {
  @Input({ required: true }) stocks: string[] = [];
  @Input() displayActions: boolean = true;
  @Input() buyHandler: (stock: string) => void = () => {};
  @Input() sellHandler: (stock: string) => void = () => {};
  @Input() removeFromWatchlistHandler: (stock: string) => void = () => {};

  activated(stock: string): boolean {
    return false;
  }
}
