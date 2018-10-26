﻿using System;
using System.Collections.Generic;
using PlatformCI.MsBuildService.Models.Interfaces;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    public class PathParser
    {
        private const string GoBack = "..";
        private const string Stay = ".";
        private const string Server = "$";
        private const string ShebangHalf = "!";
        private static readonly string[] PathSplitterChar = {"/"};

        public string GetLastItemFromPath(string anyPath)
        {
            return ParseToActionSeries(anyPath).GetLastItem();
        }

        public IPathActionSeries ParseToActionSeries(string anyPath)
        {
            var pathActionList = new List<ISourceControlPathAction>();
            var normalizedPathString = anyPath.NormalizePath();
            var splitPath = normalizedPathString.Split(PathSplitterChar, StringSplitOptions.RemoveEmptyEntries);
            foreach (var pathPart in splitPath)
                switch (pathPart)
                {
                    case Stay:
                    case Server:
                    case ShebangHalf:
                        pathActionList.Add(new PathActionStay());
                        break;
                    case GoBack:
                        pathActionList.Add(new PathActionGoBack());
                        break;
                    default:
                        pathActionList.Add(new PathActionGoAhead(pathPart));
                        break;
                }
            return new PathActionSeries(pathActionList);
        }
    }
}