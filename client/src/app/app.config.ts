import { ApplicationConfig, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideRouter, TitleStrategy } from '@angular/router';
import { MAT_CARD_CONFIG } from '@angular/material/card';
import { MAT_FORM_FIELD_DEFAULT_OPTIONS } from '@angular/material/form-field';

import { routes } from './app.routes';
import { AppTitleStrategy } from '../services';
import { provideClientHydration, withEventReplay } from '@angular/platform-browser';
import { DESIGN_SYSTEM } from '../config';
import { IDesignSystem } from '../config/design-system';

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideRouter(routes),
    provideClientHydration(withEventReplay()),
    {
      provide: TitleStrategy,
      useClass: AppTitleStrategy,
    },
    {
      provide: MAT_FORM_FIELD_DEFAULT_OPTIONS,
      useFactory: (ds: IDesignSystem) => ds.surface.MatFormFieldAppearance,
      deps: [DESIGN_SYSTEM],
    },
    {
      provide: MAT_CARD_CONFIG,
      useFactory: (ds: IDesignSystem) => ds.surface.MatCardAppearance,
      deps: [DESIGN_SYSTEM],
    },
  ],
};
