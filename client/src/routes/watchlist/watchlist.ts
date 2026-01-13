import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MatSidenavModule } from '@angular/material/sidenav';
import { Search, WatchlistTabs } from '../../components';

@Component({
  selector: 'watchlist',
  templateUrl: 'watchlist.html',
  imports: [RouterOutlet, MatSidenavModule, Search, WatchlistTabs],
})
export class Watchlist {}
