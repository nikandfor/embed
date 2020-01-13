// Generated by embed. DO NOT EDIT.

package tests

import (
	"net/http"
	"time"
)

func SecondFS(index bool) http.FileSystem {
	return fs{
		m:     mSecond,
		index: index,
	}
}

var mSecond = map[string]*file{
	`.`: {
		path:    `.`,
		size:    4096,
		mode:    020000000775,
		modTime: mustTime(time.Parse(timeFormat, "2020-01-13 12:21:12.52867937 +0300 MSK")),
		isDir:   true,
		files:   []string{`LICENSE`, `README.md`, `empty`, `go.mod`, `go.sum`, `main.go`, `tests`},
	},
	`LICENSE`: {
		path:    `LICENSE`,
		size:    1073,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2019-12-28 05:17:16.439089755 +0300 MSK")),
		content: `sQjw101JVCBMaWNlbnNlCgpDb3B5cmlnaHQgKGMpIDIwMTkgTmlraWZvciBTZXJ5YWtvdgoKUGVybWlzc2lvbiBpcyBoZXJlYnkgZ3JhbnRlZCwgZnJlZSBvZiBjaGFyZ2UsIHRvIGFueSBwZXJzb24gb2J0YWluaW5nIGEgY29weQpvZiB0aGlzIHNvZnR3YXJlIGFuZCBhc3NvY2lhdGVkIGRvY3VtZW50YXRpb24gZmlsZXMgKHRoZSAiU29mdHdhcmUiKSwgdG8gZGVhbAppbiB0aGUgU29mdAEbeCB3aXRob3V0IHJlc3RyaWN0aW9uLCBpbmNsdWRpbmcVHxRsaW1pdGEBHhAgdGhlICUCJHMKdG8gdXNlLCABqBgsIG1vZGlmAQgEZXIB1yhwdWJsaXNoLCBkaQFbGGJ1dGUsIHMBFSVLWCwgYW5kL29yIHNlbGwKY29waWVzIG9mBV8MU29mdAGeBSQhIABwIVEAdAEHDHNvbnMBEgx3aG9tMiwALCBpcwpmdXJuaXNoZQU0EGRvIHNvBXIMamVjdAE0AboUZm9sbG93AdsQY29uZGkB0jBzOgoKVGhlIGFib3ZlBcwF3xggbm90aWNlAacAICWBBYUp1g0bHHNoYWxsIGJlLTEYZWQgaW4gYR3PAHIB8SxzdGFudGlhbCBwb3IFcD7nAPB5LgoKVEhFIFNPRlRXQVJFIElTIFBST1ZJREVEICJBUyBJUyIsIFdJVEhPVVQgV0FSUkFOVFkgT0YgQU5ZIEtJTkQsIEVYUFJFU1MgT1IKSU1QTElFRCwgSU5DTFVESU5HIEJVVCBOT1QgTElNSVRFRCBUTyBUSEUgV0EFSwhJRVMBTVBNRVJDSEFOVEFCSUxJVFksCkZJVE4BVPBpRk9SIEEgUEFSVElDVUxBUiBQVVJQT1NFIEFORCBOT05JTkZSSU5HRU1FTlQuIElOIE5PIEVWRU5UIFNIQUxMIFRIRQpBVVRIT1JTIE9SIENPUFlSSUdIVCBIT0xERVJTIEJFIExJQUJMRQlrOE5ZIENMQUlNLCBEQU1BRwGeJFIgT1RIRVIKTEkRngwgV0hFARMBcCBBTiBBQ1RJT04BxzhDT05UUkFDVCwgVE9SVCARO0xXSVNFLCBBUklTSU5HIEZST00sCiFPBE9GASMESU4BOARORQlGIWsBtjmSBE9SBRAIVVNFAQsFjCAgREVBTElOR1MBPTRUSEUKU09GVFdBUkUuCg==`,
	},
	`README.md`: {
		path:    `README.md`,
		size:    1087,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2019-12-28 05:37:08.663198501 +0300 MSK")),
		content: `vwjwSVshW0RvY3VtZW50YXRpb25dKGh0dHBzOi8vZ29kb2Mub3JnL2dpdGh1Yi5jb20vbmlrYW5kZm9yL2VtYmVkP3N0YXR1cy5zdmcpCTqaOQAwKQpbIVtCdWlsZCBTdAFDCT4BeCR0cmF2aXMtY2kuSnEASC5zdmc/YnJhbmNoPW1hc3RlcimaOgAFchhjb2RlY292GTMNERguaW8vZ2gvOt4AAC8JaQAvCWksL2dyYXBoL2JhZGdlAYcdeXJGAAV5HEdvbGFuZ0NJGXoAZwUSDeYUYmFkZ2VzamkBOnEAMj0AAHJqOAANfSwgUmVwb3J0IENhcmQugwAAcgUXDGNhcmQlbRBiYWRnZW5OAGo8AAlLbj0ADAoKIyBFNgAKBQbwVSBmaWxlcyBvciBmb2xkZXJzIGludG8gZ28gZXhlY3V0YWJsZSBhbmQgZ2V0IGBodHRwLkZpbGVTeXN0ZW1gCgojIHVzYWdlCgpmaWxlCmBgYGJhc2gKCVscLS1wa2cgbXkBBkwtLXNyYyBteS50bXBsIC0tZHN0IAELUC5nbyAtLW5hbWUgTXlUZW1wbGF0ZQFIAQQkZ28KZGF0YSA6PR0dPFJlYWRBbGwoKQp0LCBlcnIBHgB0DTlULk5ldygiIikuUGFyc2Uoc3RyaW5nKAFEBCkpAVEECgoJ5gEMBGJhOqkAYUYEaWMNqgBiYRYQL2Rpc3QBxxByZWZpeDYUAAhkc3QNMgAvCTkJyihza2lwLWhpZGRlbgFlAQQ8Z28KYWxsb3dEaXJJbmRleAWxFHJ1ZQpmcwELDT4IRlMoMiYABCkKgQM0LkhhbmRsZSgiLyIsIGg1fjhlcnZlcihmcykpCmBgYAo=`,
	},
	`empty`: {
		path:    `empty`,
		size:    0,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2019-12-28 05:17:16.45108922 +0300 MSK")),
	},
	`go.mod`: {
		path:    `go.mod`,
		size:    324,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2020-01-13 12:17:51.341587026 +0300 MSK")),
		content: `xALgbW9kdWxlIGdpdGh1Yi5jb20vbmlrYW5kZm9yL2VtYmVkCgpnbyAxLjEzCgpyZXF1aXJlICgKCWdpFTBMZ29sYW5nL3NuYXBweSB2MC4wLjEyIQAZUQhjbGkJIXgwLTIwMTkxMjA5MjIxMjM3LWFkNmM5N2Y1YWZhYwoJHV4ZPQxqc29uNj4AgDI2MTgxODM4LWFjYzVmYzI0NzMwZiAvLyBpbmRpcmVjdFpKAAx0bG9nAUoIMy4wMiIAJHBrZy9lcnJvcnMBHgA4OscAZHN0cmV0Y2hyL3Rlc3RpZnkgdjEuMy4wCikK`,
	},
	`go.sum`: {
		path:    `go.sum`,
		size:    3706,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2020-01-13 12:20:36.014296441 +0300 MSK")),
		content: `+hzwcWdpdGh1Yi5jb20vYnVnZXIvanNvbnBhcnNlciB2MC4wLjAtMjAxODExMTUxOTM5NDctYmYxYzY2YmJjZTIzIGgxOkQyMUl5dXZqRENzaGoxL3FxK3BDTmQzVlpPQUVJOWp5NkJpMTMxWWxYZ0k9CmdpdOpvANgvZ28ubW9kIGgxOmJiWWxaSjdoSzF5Rng5aGY1OExQMHplWDdVaklHczIwdWZwdTNldmpyK3M9LnYAVGRhdmVjZ2gvZ28tc3BldyB2MS4xLjAdWbhKN1k4WWNXMk5paHNnbVZvL212M2xBd2wvc2tPTjRpTEhqU3NJK2M1SDM4PQpnaTU+UlkAADEBq6h2ajlqL3UxYnFudkNFZkpPd1VodGxPQVJxczMrcmtIWVkxM2pZV1RVOTdjilIA7qsADasoZ29sYW5nL2dsb2c95Fw2MDEyNjIzNTMwOC0yM2RlZjRlNmMxNGIBw6hWS3R4YWJxWFprRjI1cFk5ZWtmUkw2YTU4MlQ0UDM3LzMxWEVzdFE1cDU4MsMAtmoAHduoU0JIN3lneGk4cGZVbGFPa01NdUFRdFBJVUY4ZWNXUDVJRWwvQ1I3VlAyUU5xABxwcm90b2J1ZiHfCDMuMgHDqDZuc1BZemhxNWtSZWg2UUltSTNrNXFXek80UEVidmJJVzJjd1NmUi82eHOKUgAdq6g2bFFtNzliK2xYaU1mdmcvY1ptMFNHb2ZqSUNxVkJVdHJQNXlKTW1JQzFVTlkAFHNuYXBweSmIRS+oUWdyOXJLVzd1RFVrcmJTbVFlaURzR2E4U2pHeUNPR3R1YXNNV3d2cDJQNIJQAAAveTGgL1h4YmZtTWc4bHhlZktNN0lYQzNmQk5sLzdiUmNjNzJhQ1J6RVdybVA6UgGBECwtaXRlcmF0b3IvZ28hUwgxLjYhU6hNclV2TE1MVE14YnFGSjlremx2YXQvcllacVpuVzN1NHdrTHpXVGFGd0tzMqoAWlMAHa2oK1NkZUZCdnR5RWtYczdSRUVQMHNlVVVMcVd0YkphcExPQ1ZEYWFQRUhtVTJaAChtYWlscnUvZWFzeQFlACCZwlg5MDQwMzE5NDQxOS0xZWE0NDQ5ZGE5OC5TBKhDMXdkRkppTjk0T0pGMmI1SGJCeVFab0xkQ1dCMVlxdGcyNmc0aXJvanBjonUAVDYyNjA5MjE1OC1iMmNjYzUxOTgwMGUhPahoQjJ4bFhkSHAvcG1QWnEweTNRbm1XQUFyZHc5UHFibW90ZXhuV3gvRlU4+m4APVji4wBIb2Rlcm4tZ28vY29uY3VycmVudD1dDDgwMzCFPzQ0NC1iYWNkOWM3ZWYxZKWpqFRSTGFaOWNEL3c4UFZoOTNuc1BYYTFWclE2amx3TDVvTjhsMTRRbGNOZmc26ADWcwAd7ag2ZEpDMG1BUDRpa1lJYnZ5YzdmaWpqV0pkZFF5TG44SWczSkI1Q3FvQjlRWnoAHHJlZmxlY3QyQfQAMGmeqDlmNDEycys2Um1ZWExXWlNFelZWZ1BHSzdDMlBwaEhqNVJKcnZmeDlBV0mWVQAdsahieDJsTm5rd1ZDdXFCSXhGamZsV0pXYW5YSWIzUmxsbWJDeWx5TXJ2Z3YwMlwAMG5pa2FuZGZvci9jbGk9l1w5MTExMDE0NDEzMy1jYzJkNmMwMGRjZmYdc6hDL3QwZ2R2eHFHL0k0QUNpSWR0alhJZjcvVmVQZ3puQnZ4aFNnNlNSWjJnmnMAVDIwOTIyMTIzNy1hZDZjOTdmNWFmYWNB8qg0OWlLS3FOSE5HQUZ0SnowOFN3Tjd4WG4ydWVSNHhsVnI1WUVpVGtpYjBV8mwA/t8AMt8APkcEWDEwMzAwODA4MDctMWUyMzk1NTdlNGUwAeCoVmZ4cnUwaGJuM3dCYzBkRDhnS2dzZFpsQmgrMlNaRHdmb09mK3FXRXhNa1rgAJptAB3hqFpjbVkzcWtEMnpFdk54Ny9scWJ6UFI1TUxFUGdqU0FpRHBNTFFHMzY0Y1WedABUMjI2MTgxODM4LWFjYzVmYzI0NzMwZgHhqERBWXNZWXRINWdpK28xSUN3b2xPWW0zWE5RR1V6cGRPWU9pQWdhM00wdVn2bQD+4QAy4QAAdBrnCAQzLiWmqEU4UzlmR0JUQkk3dGJoTm9JRitEdHcxVjlscXFkdkllSGRXOXV4NzAxUndaxQAdUR2pqGE3Vmw1MENORmc4WUFUZm5ZTjV5anhpYmJybWEyR0hoUHRFYVZGbk90b0kyWAAkcGtnL2Vycm9yc2G6ADiJZqhpVVJVclJHeFBVTlBkeTUvSFJTbStZajZva0o2VXRMSU5OMFE5TTQraDNJdk0AHaGoYndhd3hmSEJGTlYrTDJoVXAxckhBRHVmVjNJTXRuRFJkZjFyNU5JTkVsMDZUABRtZXphcmQOMgskZGlmZmxpYiB2MQ4ZDEETqDREQndERTBOR3lRb0JIYkxRWVB3U1VQb0NNV1I1QkV6SWsvZjFsWmJBUU2WVQAdsahpS0g3N2tvRmhZeFRLMXBjUm5rS2txZlRvZ3NiZzdnWk5WWTRzUkRZWi80MlwAMHJha3lsbC9zdGF0aWshVQQxLi5cCKhPRWk5d0pWL2ZNVUFHeDFlTmpxNzVES0RzSlZ1RXYxVTBvWWRYNkdYOFpzMlcAMHN0cmV0Y2hyL29iangJVy6uAKhIRmtZOTE2SUYrcndkRGZNQWtWN090d3VxQlZ6ckU4R1I2R0Z4K3dFeE1FdlcAADIyVwCocXQwOVlhOHZhd0x0ZTZTTm1UZ0NzQVZ0WXRhS3pFY244QVRVb0hNa0VxRVZXABh0ZXN0aWZ5IbRtAqhUaXZDbi9wZUJRN1VZOG9vSWNQZ1pGcFROU3owUTJVNlVyRmxVZnFiZTBRjlMAPbKwTTVXSXk5RGgyMUlFSWZuR0N3WEdjNWJaZktOSnRmSG0xVVZVZ1puKzlFST0K`,
	},
	`main.go`: {
		path:    `main.go`,
		size:    8606,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2020-01-13 12:21:12.520679724 +0300 MSK")),
		content: `nkOIcGFja2FnZSBtYWluCgovL2dvOmdlbmVyYXRlIGdvIHJ1biABG3AuZ28gLS1wa2cgdGVzdHMgLS1za2lwLWhpZGRlbgEOJHJjIC4gLS1kc3QJIgwvZGlyATUFL0INAAwyLmdv/m0AOm0AAVM0IC0tbmFtZSBTZWNvbmT+fABifAAUZ28ubW9kMu4ADGZpbGXu1AAF1BBlbXB0eTJXAAUSIUcN2QBFASKICgppbXBvcnQgKAoJImVuY29kaW5nL2Jhc2U2NCIKCSJmbXQBBwRpbwkGGC9pb3V0aWwBDQRvcwEGDHBhdGgRCAW5EREUc3RyaW5nBSQwdGV4dC90ZW1wbGF0ZQEtgHRpbWUiCgoJImdpdGh1Yi5jb20vZ29sYW5nL3NuYXBweQElHRwwbmlrYW5kZm9yL2NsaWIcAAx0bG9nOh0ATHBrZy9lcnJvcnMiCikKCnR5cGUgIV5EIHN0cnVjdCB7CglQYXRoICAgARIgaW5nCglTaXplARA4aW50NjQKCU1vZFRpbWUgAbMALgEKLAoJRmlsZXMgICBbXQnkIAoJQ29udGVudBVECE1vZAVECG9zLgEtAQ9UCglJc0RpciAgIGJvb2wKfQoKZnVuY2UqBCgpAYgkY2xpLkFwcCA9IAEK6ENvbW1hbmR7CgkJTmFtZTogICAiZW1iZWQiLAoJCVVzYWdlOiAgIltPUFRJT05TXSB7PGV4Y2x1ZGVfIYEEPn0FKBxBY3Rpb246IAU6ATkkRmxhZ3M6IFtdKgFlAQ4BYgV8CE5ldwEQQCgic3JjLHMiLCAic3RhdGljAQogZm9sZGVyIG9mKT0EdG8JUgQiKQFUNj4AEGRzdCxkATQJPgAvBXsUZGVkLmdvARYUb3V0cHV0BUdOPgCNQRQscGtnLHABMAlGAQoNGwAgYZROOwABGAQsbgEqAQQocHJlZml4IGFsbCAhaCEHLHMgd2l0aCBnaXZlbmJGAJ2JJCxoIiwgZmFsc2UhDAxraXAgjaFBQRxzIGluIHNyY06KABRmb3JjZS1JYgBzHUQIYWRkQSJBeBRkZWZpbmkJlgBlAZBmmwAMSGVscCGZIWsYfSwKCX0KCiVySFJ1bkFuZEV4aXQob3MuQXJncylVViV1CChjICXmTUsEKSBlCUFrIVUALAEPCCA6PUGkHFN0YXQoYy5TZbJFARgiKSkKCWlmBSQUIT0gbmlsATgYCXJldHVybgEWAYgBIQAhAUwALkXbSccVJEGtQf0IKGMpAS0QCgl2YXItOwhbXSohRBQKCWRpcnMBhAxtYXBbaUwJGgB7AeUQZXJyID0FMkHSEC5XYWxrPqUAACwl3QQocCFXIGluZywgaW5mbwHSAXkMSW5mbwnmBfk9AAG5EHAgPT0gBVkQQ2xlYW4dWhBkc3QiKWWTBAkJgVMcLlByaW50ZihFHDggZHN0OiAldiIsIHApCgkV70RuaWwKCQl9CgkJZm9yIF8sIGEB2BhyYW5nZSBjJZ4lQRV6IQAMcGF0aA1+AGENcEpxAAhhcmcycQAxYQlyBXMhLhAJcmVhbIXMEDo9IHAKCWoAcAENOtkBDDsgcHAh1BAiIiAmJhEMBC4iFZ8piBRzLkhhc1BlRRQocCwgcHARoYGhKHBbbGVuKHBwKTpdIQMUfSBlbHNlCUERnhBmbXQuRaF5BGYoxSNEICV2IGRvZXMgbm90IGhhdmUgbaIpSwViBUw9RlKKABAiLi4vIg2NCYwAMwWGATkEaWYp1xAiIiB8fA0LAC4R2wEtCCIuIg1gASsIWzBdQQUIJy8nCbwBIghwWzENTylDFGMuQm9vbCn8GqEIEX4AbiFdJcQQQmFzZShJDwxpZiBuJV4ALiVfAG4RZAAuDWRK4QGJVi3kMcYlDQFOQdIyYANZdFVEGFNraXBEaXIhQCUINYY+JAIEaWYVYQGlAHAJ+DLlAgHRICAlcSBpc2RpciG2DCglcSlB7mldEYsALBVJAaouAgJyTgBGSQAFPyFtAGYhTwAmAdTFk0HLwfYEIHCB+eWWBQ8lEOGkACjJfeFcGRkBDhkZ4a8AOgky7ccNHIWwATYFTgUOCRoJiuXgID0gYXBwZW5kKAUPDCwgZikBpgRpZjI0AClBgccIW3Bdga5BX30PDGRhdGGJhAg6PSAWOwkMLlJlYakZMuwBRm0FNcaFv0BzLldyYXAoZXJyLCAicmVhZMF+RQMAIiFDCboBRmGbAXcAKSHiADAJoQQJegGAFlsJMC5FbmNvZGUobmlsLCAFKUU4hbAW9wkMLlN0ZAUlCGluZw0uBFRvzR0AeiUGDAkJZi4e9AgIPSBhAUFhgCU+GGlmIGQsIGZ5HiBTcGxpdChwKTsBFQQhPWGFCZYBKmGBACINEAgJZCAFHhFPASEAbAG5BdMkKSAtIDE7IGRbbDa2AwU2DGRbOmyJjAE9RYMgVigiZGlycyIptcBB0UHbIdEBCAAgASBteAWzACwFEiUQARsIW2RdxUkAcwHkSQwyFwAB5gEtBfftG2VyAH16PQcAbzbOAQx3YWxrIc4hyQFM4SkkdyBpby5Xcml0ZeEFBG9zIZQIb3V05XQAcSEUudbFsAg7IHElXAAtKUwAZlV1KG9zLkNyZWF0ZShxAaxKZwINukZmAgBjBTjB8gGaJeI0ZGVmZXIgZi5DbG9zZShFDAB3SfcBuQB0AZke6AsOUQoEKCIBOhx0LkZ1bmNzKBUaAQ8ITWFwYegIInNw5V4EIjrFMwBTCQ4O5wgUKQoJdCA9GU9ETXVzdCh0LlBhcnNlKGAvLyBHGloOIGQgYnkge3sgLhprDlBvciB9fS4gRE8gTk9UIEVESVQuCgoePgoBKBEMOH19Cgp7eyBpZiBhbmQgLg5BCgQgKMGgAC4SwgkAUxIjDBRzKSAtfX0uJQ0cbmV0L2h0dHAOawwOFgwOWgwIe3stafY6LAAMYnl0ZRIDDUpaDSlxAUYWXQ0uWAA2XA0Mc3luYwEoCW5yNw0e5QwIKAoJRecWhgplcKXaKukMBAlzNuoMBAltRusMARRGygwECWkuywwO9AilMi4SDQQJYzYTDRQKCQlvbmMFSwHOJC5PbmNlCgkJZGWBjAxkIFtdISoWKAoEZnMutgAAbQGzQiEK4Z0MbmRleBZNDUV9BGZzAYYuOwAEciABVQRzLqFxAGXhHhp2CggJZnMBew7iDQE7CCkKCmFyIV0QRm9ybWFBdlQiMjAwNi0wMS0wMiAxNTowNDowNS45EQEoIC0wNzAwIE1TVCIJPBQoCglFcnJlFQBkBbAAPXFPFE5ldygiY2EuAGSFeGEQIEVyck91dE9mUhIeCjYrAAhvdXQOZw0SOQoMIikKKUlaBG5kRa0OygxFwkWrEH19RlMoOQYEKSBBlSguRmlsZVN5c3RlbSH1bdwEZnNhbABtwfIIICBtYQQNQuFcBUPhJwhkZXhheQ5GDGkXBG9yNgsDbR4BPGkWBZEkKGZzIGZzKSBPcA4VCiGBGGluZykgKF8ZkYmORfHlMwRpZqXUAHDNpA6ICh4pCRqNCSqMCQB9hfUSsAvNKgEdEssJBH0KgeoEb2uBeBhmcy5tW3BdBS0IIW9rAf4R/8XrBG9zDpEKGE5vdEV4aXMl+QwKCWYuQZUMLkRvKAHKGsEMEGlmIGYuUb8EPT0RdQ1NAAqhCQQKCaWlAHpVuAQJegnjAD1OQwcAREXl7UEVW3KOBRFmBGYubSIRX+3ECVPpxAB6DmANanIGQeQICglmGmAJZSYAe2ErKDogZiwgZnM6IGZzIUgkZi5yLlJlc2V0KBV1oeBRQgRmLCFHQRwhKxAgKGYgKglMBCkgRfgAKCJgDQggZi7hyQwtMTsgDT0OPw4e8RAuOQAYU2VlayhvZg5TCRx0NjQsIHdoZYEkFGludCkgKBKAEQX3AG9VNQVaDD0gLTEu3gEIMCwgdX0B7BWzCC5yLhFiEVwmIw8uiACBIQQocC3NDCkgKG4BgwV7NrQCwn8ACVlmdQAEQXQVdwQsIDUJ/oIAPoIAAWYFX1aJABL+DwQpIA5GEIE9IjcPLXctxARmLkE0ZgUCAesIZGlyKTwQKSAoW10a/RI+UAAgaWYgIWYuZnMuhXguyQFlpxZBCRovEAxyZXMgMkkABAoJDl8PIYgAPIVCCZYEcymBQw4mCgByAQwIPCBuDVcqugkccmVzLCBtW3AOdA8YSm9pbihmLgEMACwN2RxzW2YuZF0pXWGWEGYuZCsrQUqFTAFtBD09BV4VbqFOLRgFWBRpby5FT0YJMy4YAAHNMvkCEGlsZSkgDtETBCgprSnFLwggIHtxLAGVFlEOCaEIKSB9fXIFQBb+DEGkBDY0DT8ZQARmLuXXQjUAFhoNoQYIRmlsEpIUMjUADG1vZGVSNQAORA0EKCnlHhL/FC5qAB4xCEJtABpSDeWeSqIAFj8IQjYADFN5cygF1xxlcmZhY2V7fS4YAS5EBAxtdXN0BZgAdBmYMq8D4cwNriXEDksSKscSEHBhbmljDksNoTqVMAB0IasIe3sg4W4OxwkOFQgAbf0sBCA9QocIAHvln+WtACBlIABz4agYCWAgKyAiYAE7DsYOJCB9fWAiICsgYDpFOCHh5YlaJgDhkyHDBSQBRCHvBCB95ashphUYIdIIIHwgGkELFCAiMCVvIhkoISAAOj0qCGltZRo8CwELFukICCwgIgFhGhwPECB9fSIpElUPDs4KCGlmIBYZDwXSFvIJAXsBMSXfDWtBkwBzARkAWxoOF2EcAUEpFhQkaSwgJGXhoAAuQXghIBZcCAQkaQFBAU8hcAR9fTkOAQ0lMC2MAH0J/AFUCSsECgkFDRIzFSKHDgB9ARfttgA6PVwVTgAsFU8ACg5IFR1PCXwEXG4O9RQAChp3FCx0LkV4ZWN1dGUodywu9wEEaW5VngHpACIidAwAIiEELqQXHrEWCRkAICqxDRqADBY7FhK4FhEkMicADmcJDSQlZRklBQ9BKBZrFhqpDAQiOiKuEhLADCKCFiHUbiEIYR86kw4eqw0OrwqhjY3VLtAEKuEVXmgWDl4QaZhK/hAmwQ4SRxdqwQhKoAAeARERoGGsIf2tVWX1wRABghplCiGvfgcRBUNSxAkSEBEyBxElGAF4/qsP/qsP7qsPGqsPFlwPGqEbFrQMDjgR/msPNmsPKmgcEpgNACJlhIFoIXcOPxthxE4qDxYgDwBpJn0ccgMPYegOkBYFTgE7AREEbmSl1AApGncbgcZhYAggfX0OuQgAZQ69GARpbxKwCQBlDiMZgR4+mgAa4wsAOlIgAgBELuQLAGMdX1LYAj6BBnUlWt8LRrYRPkIAFpkZJVghwA3VAHIOQA4WmRAhDlomAAhuaWwJKC0lHrIIHcYBUBBBbGwoKRolCwAgyfo+IwFhbxXMKu8UATs6TgAhc5IcAQnaAVolyTHaDRiBkj3gAH0WkA8ugQEEID3dDllqAGDu8AWp8CqRHbXUwQaZ2mbqBRkkQYelowQJIg2eCUUIb250Do4PUiQBDf0SUQ8yGRZxxLEYMHJldHVybiBuaWwKfQo=`,
	},
	`tests`: {
		path:    `tests`,
		size:    4096,
		mode:    020000000775,
		modTime: mustTime(time.Parse(timeFormat, "2020-01-13 12:20:26.938698338 +0300 MSK")),
		isDir:   true,
		files:   []string{`dir_test.go`, `empty.go`, `empty_test.go`, `file.go`, `file_test.go`},
	},
	`tests/dir_test.go`: {
		path:    `tests/dir_test.go`,
		size:    2848,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2020-01-13 12:20:06.899585701 +0300 MSK")),
		content: `oBZwcGFja2FnZSB0ZXN0cwoKaW1wb3J0ICgKCSJpbyIFBhgvaW91dGlsAQ0Eb3MBBgErCGluZwULbGltZSIKCgkiZ2l0aHViLmNvbS9zdHJldGNoci8FKJBmeS9hc3NlcnQiCikKCmZ1bmMgVGVzdERpck9wZW4odCAqdGVzAU98LlQpIHsKCWZzMSA6PSBGUyhmYWxzZSkKCglmLCBlcnIBFgxmczEuBTcFhBRzL2Rpcl8BZhguZ28iKQoJCWkkLk5vRXJyb3IodAU4AUIAcwUKBUIYLlN0YXQoKWotAA1GEEVxdWFsAUQAIi5mACQsIHMuTmFtZSgpAVgBlyA9IGYuQ2xvc2VyVQAELy8VyAA9HccUZ28ubW9kjrwAAWfGuwANYC62AELgACRpbnQ2NCgzMjQpAScIU2l6Bd0RJwxUcnVlIQZAcy5Nb2RlKCkmMDY0NCA9PSABCDlHESdBETAuTm93KCkuQWZ0ZXIoBTgEVGklMhkvNV0AZiHlAXgQSXNEaXIhVhF5BE5pJX4Qcy5TeXMBGBQKCWRhdGE1yUmVJC5SZWFkQWxsKGYZXABOQv8BEGZjb250If4cYG1vZHVsZSBdroBuaWthbmRmb3IvZW1iZWQKCmdvIDEuMTMKCnJlcXVpcmVhDB0wUGdvbGFuZy9zbmFwcHkgdjAuMC4xCi4hABlRCGNsaQkhcDAtMjAxOTEyMDkyMjEyMzctYWQ2Yzk3ZjVhZmFjWj0ADGpzb242PgCAMjYxODE4MzgtYWNjNWZjMjQ3MzBmIC8vIGluZGlyZWN0WkoADHRsb2cBSggzLjAyIgAQcGtnL2VhKQBzAR4AODrHAD7GAwggdjEFQggpCmA1oznGIWEgLCBzdHJpbmcoIahlHwhvZmY1sCxmLlNlZWsoMCwgaW8FCxhDdXJyZW50ZrgBUQkVZkmpDGxlbigFcBgpKSwgb2ZmYYgAbh1nCChpb0UWEGVyQXQpBQoEQXQFkyBbOjQwXSwgMTCqcwAQNDAsIG4BYqbqAwBfCYxlsDbyAAxTdGFyHfAV2AhFcnKFQQBkBToBbh1EQeEMKG5pbBm1ljgAUgcBEG5pbCwgPQBiSwAAfTarBQGFCGRpck6uBWFXLFNlY29uZEZTKHRydUKyBQAuqbGO5AQAOiU1mqAFAFSVmQBJPjsEFdcEIi4y+wQECgkykQIR1p3HPjwEEAlpZiAhzREATMGvBCwgAUMANMWrXAlyZXR1cm4KCX0KCgl2YXIgbGlzdCBbXWkHJAoJZm9yIF8sIGYB2BxyYW5nZSBmZsHjAAkFKCA9IGFwcGVuZCgBDggsIGbVVwVPAWZRdBGoADJZBJqoAAAyRqgA/pUAapUANXcUaW8uRU9GTc5aQwEAMf6bADabAC2kMEVsZW1lbnRzTWF0Y2jBtyldMZRUeyJMSUNFTlNFIiwgIlJFQURNRS5tZAENEGVtcHR5AQn1MgEKCHN1bQEUDG1haW4OYwgELCAWeAgEIn1hpeXTADIF5iHiAGZR6gAvdusCHTIAMqLSB1Xt4XoEczIRf5koEX0kbm90LWV4aXN0cx2GET4gb3MuSXNOb3RFASEUKGVycikpWuoDBE5vguwDAGbhugGJAGZGAwEdfD6qAgAK0o0JAFT+7QN17W5FA01FME5pbCh0LCBmZikKfQo=`,
	},
	`tests/empty.go`: {
		path:    `tests/empty.go`,
		size:    202,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2020-01-13 12:18:01.529136104 +0300 MSK")),
		content: `ygHwdS8vIEdlbmVyYXRlZCBieSBlbWJlZC4gRE8gTk9UIEVESVQuCgpwYWNrYWdlIHRlc3RzCgppbXBvcnQgKAoJImJ5dGVzIgoJImlvIgopCgpmdW5jIEVtcHR5UmVhZGVyKCkgaW8uUmVhZGVyIHsKCXJldHVybiAFNwwuTmV3DSYUbmlsKQp9Pj8AHEFsbCgpIFtdAS4YIHsKCXJldAE9CG5pbAEsPHZhciBjRW1wdHkgPSBgYAo=`,
	},
	`tests/empty_test.go`: {
		path:    `tests/empty_test.go`,
		size:    231,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2019-12-28 05:17:16.45108922 +0300 MSK")),
		content: `5wHwQHBhY2thZ2UgdGVzdHMKCmltcG9ydCAoCgkiaW8iCgkidGVzdGluZyIKCgkiZ2l0aHViLmNvbS9zdHJldGNoci90ASB8ZnkvYXNzZXJ0IgopCgpmdW5jIFRlc3RFbXB0eSh0ICoNRSwuVCkgewoJZCA6PSAFHChSZWFkQWxsKCkKCQlABC5FCTMgLCBkKQoKCXIgLioABGVyASkUXywgZXJyARkEci4BQAwobmlsHUBYcXVhbCh0LCBpby5FT0YsIGVycikKfQo=`,
	},
	`tests/file.go`: {
		path:    `tests/file.go`,
		size:    776,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2020-01-13 12:18:01.217149913 +0300 MSK")),
		content: `iAbwXi8vIEdlbmVyYXRlZCBieSBlbWJlZC4gRE8gTk9UIEVESVQuCgpwYWNrYWdlIHRlc3RzCgppbXBvcnQgKAoJImJ5dGVzIgoJImVuY29kaW5nL2Jhc2U2NCIKCSJpbyIKAQboL2lvdXRpbCIKCgkiZ2l0aHViLmNvbS9nb2xhbmcvc25hcHB5IgopCgpmdW5jIFJlYWRlcigpIGlvLlIFDDwgewoJeiwgZXJyIDo9IGJhAWAULlN0ZEVuCXNMLkRlY29kZVN0cmluZyhjKQoJaWYFLhQhPSBuaWwBP0AJcGFuaWMoZXJyKQoJfQoJchVPCX8NQxgobmlsLCB6kkIAFGV0dXJuICUAEC5OZXdSBasQKHIpCn0dyRxBbGwoKSBbXQEnAYgMZGF0YRV7KRcF5QEoDUIAKap+AAFPAXD0QgF2YXIgYyA9IGB4QUxnYlc5a2RXeGxJR2RwZEdoMVlpNWpiMjB2Ym1scllXNWtabTl5TDJWdFltVmtDZ3BuYnlBeExqRXpDZ3B5WlhGMWFYSmxJQ2dLQ1dkcEZUQk1aMjlzWVc1bkwzTnVZWEJ3ZVNCMk1DNHdMakV5SVFBWlVRaGpiR2tKSVhnd0xUSXdNVGt4TWpBNU1qSXhNak0zTFdGa05tTTVOMlkxWVdaaFl3b0pIVjRaUFF4cWMyOXVOajRBZ0RJMk1UZ3hPRE00TFdGall6Vm1ZekkwTnpNd1ppQXZMeUJwYm1ScGNtVmpkRnBLQUF4MGJHOW5BVW9JTXk0d01pSUFKSEJyWnk5bGNuSnZjbk1CSGdBNE9zY0FaSE4wY21WMFkyaHlMM1JsYzNScFpua2dkakV1TXk0d0Npa0tgCg==`,
	},
	`tests/file_test.go`: {
		path:    `tests/file_test.go`,
		size:    481,
		mode:    0664,
		modTime: mustTime(time.Parse(timeFormat, "2020-01-13 12:20:26.926698869 +0300 MSK")),
		content: `4QNkcGFja2FnZSB0ZXN0cwoKaW1wb3J0ICgKCSIBEmxpbmciCgoJImdpdGh1Yi5jb20vc3RyZXRjaHIvBSB4ZnkvYXNzZXJ0IgopCgpmdW5jIFRlc3RGaWxlKHQgKgUkYG5nLlQpIHsKCWQgOj0gUmVhZEFsbCgpCgkJOkQuRXF1YWwodCwgYG1vZHVsZSAdboBuaWthbmRmb3IvZW1iZWQKCmdvIDEuMTMKCnJlcXVpcmUBqx0wUGdvbGFuZy9zbmFwcHkgdjAuMC4xCi4hABlRCGNsaQkhdDAtMjAxOTEyMDkyMjEyMzctYWQ2Yzk3ZjVhZmFjClY9AAxqc29uNj4AgDI2MTgxODM4LWFjYzVmYzI0NzMwZiAvLyBpbmRpcmVjdFpKAAx0bG9nAUoIMy4wMiIAJHBrZy9lcnJvcnMBHgA4OscAPoYBCCB2MQVCRCkKYCwgc3RyaW5nKGQpKQp9Cg==`,
	},
}
