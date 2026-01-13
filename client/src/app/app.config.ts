import {
  inject,
  ApplicationConfig,
  provideAppInitializer,
  provideBrowserGlobalErrorListeners,
  PLATFORM_ID,
} from '@angular/core';
import { provideRouter, TitleStrategy } from '@angular/router';
import { MAT_CARD_CONFIG } from '@angular/material/card';
import { MAT_FORM_FIELD_DEFAULT_OPTIONS } from '@angular/material/form-field';

import { routes } from './app.routes';
import { AppTitleStrategy } from '../services';
import { provideClientHydration, withEventReplay } from '@angular/platform-browser';
import { DESIGN_SYSTEM } from '../config';
import { IDesignSystem } from '../types/design-system.types';
import { applyDesignSystem } from '../config/design-system.apply';
import { provideHttpClient, withFetch } from '@angular/common/http';

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideRouter(routes),
    provideClientHydration(withEventReplay()),
    provideHttpClient(withFetch()),
    {
      provide: TitleStrategy,
      useClass: AppTitleStrategy,
    },
    provideAppInitializer(() => {
      const ds = inject(DESIGN_SYSTEM);
      const platform_id = inject(PLATFORM_ID);
      applyDesignSystem(ds, platform_id);
    }),
    {
      provide: MAT_FORM_FIELD_DEFAULT_OPTIONS,
      useFactory: (ds: IDesignSystem) => ({
        appearance: ds.surface.MatFormFieldAppearance,
      }),
      deps: [DESIGN_SYSTEM],
    },
    {
      provide: MAT_CARD_CONFIG,
      useFactory: (ds: IDesignSystem) => ({ appearance: ds.surface.MatCardAppearance }),
      deps: [DESIGN_SYSTEM],
    },
  ],
};
