import { Component, inject } from '@angular/core';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { DESIGN_SYSTEM } from '../../config';
import { Router, RouterLink, RouterOutlet } from '@angular/router';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { TitleCasePipe } from '@angular/common';

@Component({
  selector: 'dashboard',
  templateUrl: 'dashboard.html',
  imports: [
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    RouterLink,
    MatSidenavModule,
    RouterOutlet,
    MatListModule,
    TitleCasePipe,
  ],
})
export class Dashboard {
  readonly ds = inject(DESIGN_SYSTEM);
  readonly title = this.ds.app.APP_NAME;
  readonly router = inject(Router);
  readonly isDarkMode = window.matchMedia('(prefers-color-scheme: dark)').matches;
  opened = true;
  navItems = [
    { label: 'watchlist', icon: 'bookmark' },
    { label: 'profile', icon: 'person' },
  ];
  activeLink = this.router.url.split('/').pop();
  toggleDarkMode() {
    this.ds.color.theme =
      this.ds.color.theme === 'system'
        ? this.isDarkMode
          ? 'light'
          : 'dark'
        : this.ds.color.theme === 'light'
        ? 'dark'
        : 'light';
    document.documentElement.style.setProperty('--theme', this.ds.color.theme);
  }
}
