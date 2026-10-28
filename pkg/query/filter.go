package query

import (
	"github.com/leonelquinteros/gotext"

	"github.com/Jguer/yay/v13/pkg/settings/parser"
	"github.com/Jguer/yay/v13/pkg/text"
)

func RemoveInvalidTargets(logger *text.Logger, targets []string, mode parser.TargetMode) []string {
	filteredTargets := make([]string, 0)

	for _, target := range targets {
		dbName, _ := text.SplitDBFromName(target)

		if dbName == "aur" && !mode.AtLeastAUR() {
			logger.Warnln(gotext.Get("%s: can't use target with option --repo -- skipping", text.Cyan(target)))
			continue
		}

		if dbName != "aur" && dbName != "" && !mode.AtLeastRepo() {
			logger.Warnln(gotext.Get("%s: can't use target with option --aur -- skipping", text.Cyan(target)))
			continue
		}

		filteredTargets = append(filteredTargets, target)
	}

	return filteredTargets
}

// QualifyAURTargets prefixes unqualified targets with the aur db when running
// in --aur mode, so repo packages that provide or replace the name are not
// picked over the AUR package.
func QualifyAURTargets(mode parser.TargetMode, targets []string) []string {
	if mode != parser.ModeAUR {
		return targets
	}

	qualified := make([]string, len(targets))

	for i, target := range targets {
		if dbName, _ := text.SplitDBFromName(target); dbName == "" {
			target = "aur/" + target
		}

		qualified[i] = target
	}

	return qualified
}
